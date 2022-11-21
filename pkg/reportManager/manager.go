// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package reportManager

import (
	"context"
	"go.uber.org/zap"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/util/workqueue"
	"time"
)

const (
	queueMaxRetries = 100
)

type reportManager struct {
	logger          *zap.Logger
	reportDir       string
	collectInterval time.Duration
	queue           workqueue.RateLimitingInterface
}

var globalReportManager *reportManager

func InitReportManager(logger *zap.Logger, reportDir string, collectInterval time.Duration) {
	if globalReportManager != nil {
		return
	}

	globalReportManager = &reportManager{
		logger:          logger,
		reportDir:       reportDir,
		collectInterval: collectInterval,
		queue:           workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "reportManager"),
	}
	go globalReportManager.runWorker()
}

func (s *reportManager) runWorker() {

	// TODO: wait for all agent grpc is ready
	s.logger.Info("waiting for all agent grpc server ready")

	s.logger.Info("all agent grpc server ready, start worker")

	//
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	defer s.queue.ShutDown()

	// please do not run more than one worker, or else it races to write reports
	go wait.UntilWithContext(ctx, s.worker, time.Second)

	// periodically trigger sync
	for {
		TriggerSyncReport("periodicallyTrigger")
		<-time.After(s.collectInterval)
	}
}

func TriggerSyncReport(triggerName string) {
	if globalReportManager != nil {
		globalReportManager.logger.Sugar().Debugf("trigger to sync agent report from source %v", triggerName)
		// s.queue.AddRateLimited(triggerName)
		globalReportManager.queue.AddAfter(triggerName, 10*time.Second)
	}
}

// --------------

func (s *reportManager) worker(ctx context.Context) {
	for s.processNextWorkItem(ctx) {
	}
}

func (s *reportManager) processNextWorkItem(ctx context.Context) bool {
	key, quit := s.queue.Get()
	if quit {
		return false
	}
	defer s.queue.Done(key)

	err := s.syncHandler(ctx, key.(string))
	if err == nil {
		s.queue.Forget(key)
	} else {
		s.logger.Sugar().Warnf("worker failed , error=%v", err)
		if s.queue.NumRequeues(key) < queueMaxRetries {
			s.queue.AddRateLimited(key)
		}
	}
	// handle nex item
	return true
}
