// Copyright (c) Abstract Machines
// SPDX-License-Identifier: Apache-2.0

package middleware

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/absmach/magistrala/re"
	"github.com/absmach/supermq/pkg/authn"
	"github.com/absmach/supermq/pkg/messaging"
)

var _ re.Service = (*loggingMiddleware)(nil)

type loggingMiddleware struct {
	logger *slog.Logger
	svc    re.Service
}

func LoggingMiddleware(svc re.Service, logger *slog.Logger) re.Service {
	return &loggingMiddleware{logger, svc}
}

func (lm *loggingMiddleware) AddRule(ctx context.Context, session authn.Session, r re.Rule) (res re.Rule, err error) {
	defer func(begin time.Time) {
		args := []any{
			slog.String("duration", time.Since(begin).String()),
			slog.String("domain_id", session.DomainID),
		}
		if err != nil {
			args = append(args, slog.String("error", err.Error()))
			lm.logger.Warn(fmt.Sprintf("Add %s rule failed", r.Name), args...)
			return
		}
		lm.logger.Info(fmt.Sprintf("Add %s rule completed successfully", r.Name), args...)
	}(time.Now())
	return lm.svc.AddRule(ctx, session, r)
}

func (lm *loggingMiddleware) ViewRule(ctx context.Context, session authn.Session, id string) (res re.Rule, err error) {
	defer func(begin time.Time) {
		args := []any{
			slog.String("duration", time.Since(begin).String()),
			slog.String("domain_id", session.DomainID),
			slog.String("rule_id", id),
		}
		if err != nil {
			args = append(args, slog.String("error", err.Error()))
			lm.logger.Warn(fmt.Sprintf("View %s rule failed", id), args...)
			return
		}
		lm.logger.Info(fmt.Sprintf("View %s rule successful", id), args...)

	}(time.Now())
	return lm.svc.ViewRule(ctx, session, id)
}

func (lm *loggingMiddleware) UpdateRule(ctx context.Context, session authn.Session, r re.Rule) (res re.Rule, err error) {
	defer func(begin time.Time) {
		args := []any{
			slog.String("duration", time.Since(begin).String()),
			slog.String("domain_id", session.DomainID),
			slog.Group("rule",
				slog.String("id", r.ID),
				slog.String("name", r.Name),
			),
		}
		if err != nil {
			args = append(args, slog.String("error", err.Error()))
			lm.logger.Warn(fmt.Sprintf("Update %s rule failed", r.ID), args...)
			return
		}
		lm.logger.Info(fmt.Sprintf("Update %s rule successful", r.ID), args...)
	}(time.Now())
	return lm.svc.UpdateRule(ctx, session, r)
}

func (lm *loggingMiddleware) ListRules(ctx context.Context, session authn.Session, pm re.PageMeta) (pg re.Page, err error) {
	defer func(begin time.Time) {
		args := []any{
			slog.String("duration", time.Since(begin).String()),
			slog.String("domain_id", session.DomainID),
			slog.Group("page",
				slog.Uint64("offset", pm.Offset),
				slog.Uint64("limit", pm.Limit),
				slog.Uint64("total", pg.Total),
			),
		}
		if err != nil {
			args = append(args, slog.String("error", err.Error()))
			lm.logger.Warn("List rules failed", args...)
			return
		}
		lm.logger.Info("List rules successful", args...)
	}(time.Now())
	return lm.svc.ListRules(ctx, session, pm)
}

func (lm *loggingMiddleware) RemoveRule(ctx context.Context, session authn.Session, id string) (err error) {
	defer func(begin time.Time) {
		args := []any{
			slog.String("duration", time.Since(begin).String()),
			slog.String("domain_id", session.DomainID),
			slog.String("rule_id", id),
		}
		if err != nil {
			args = append(args, slog.String("error", err.Error()))
			lm.logger.Warn(fmt.Sprintf("Remove %s rule failed", id), args...)
			return
		}
		lm.logger.Info(fmt.Sprintf("Remove %s rule successful", id), args...)
	}(time.Now())
	return lm.svc.RemoveRule(ctx, session, id)
}

func (lm *loggingMiddleware) EnableRule(ctx context.Context, session authn.Session, id string) (res re.Rule, err error) {
	defer func(begin time.Time) {
		args := []any{
			slog.String("duration", time.Since(begin).String()),
			slog.String("domain_id", session.DomainID),
			slog.String("rule_id", id),
		}
		if err != nil {
			args = append(args, slog.String("error", err.Error()))
			lm.logger.Warn(fmt.Sprintf("Enable %s rule failed", id), args...)
			return
		}
		lm.logger.Info(fmt.Sprintf("Enable %s rule successful", id), args...)
	}(time.Now())
	return lm.svc.EnableRule(ctx, session, id)
}

func (lm *loggingMiddleware) DisableRule(ctx context.Context, session authn.Session, id string) (res re.Rule, err error) {
	defer func(begin time.Time) {
		args := []any{
			slog.String("duration", time.Since(begin).String()),
			slog.String("domain_id", session.DomainID),
			slog.String("rule_id", id),
		}
		if err != nil {
			args = append(args, slog.String("error", err.Error()))
			lm.logger.Warn(fmt.Sprintf("Disable %s rule failed", id), args...)
			return
		}
		lm.logger.Info(fmt.Sprintf("Disable %s rule sucessful", id), args...)
	}(time.Now())
	return lm.svc.DisableRule(ctx, session, id)
}

func (lm *loggingMiddleware) StartScheduler(ctx context.Context) (err error) {
	defer func(begin time.Time) {
		args := []any{
			slog.String("duration", time.Since(begin).String()),
		}
		if err != nil {
			args = append(args, slog.String("error", err.Error()))
			lm.logger.Warn("Start scheduler failed", args...)
			return
		}
		lm.logger.Info("Start scheduler successful", args...)
	}(time.Now())
	return lm.svc.StartScheduler(ctx)
}

func (lm *loggingMiddleware) ConsumeAsync(ctx context.Context, msgs interface{}) {
	defer func(begin time.Time) {
		args := []any{
			slog.String("duration", time.Since(begin).String()),
		}
		if m, ok := msgs.(*messaging.Message); ok {
			args = append(args,
				slog.String("channel", m.Channel),
				slog.String("payload_size", fmt.Sprintf("%d", len(m.Payload))),
			)
		}
		lm.logger.Info("Message consumption completed", args...)
	}(time.Now())

	lm.svc.ConsumeAsync(ctx, msgs)
}

func (lm *loggingMiddleware) Errors() <-chan error {
	return lm.svc.Errors()
}
