-- +goose Up
CREATE INDEX idx_subscriptions_user_id ON subscriptions(user_id);
CREATE INDEX idx_subscriptions_user_service ON subscriptions(user_id, service_name);

CREATE UNIQUE INDEX uniq_active_subscription
    ON subscriptions(user_id, service_name)
    WHERE status = 'active';

CREATE INDEX idx_transactions_subscription_created
    ON subscription_transactions(subscription_id, created_at);


-- +goose Down
DROP INDEX IF EXISTS idx_subscriptions_user_id;
DROP INDEX IF EXISTS idx_subscriptions_user_service;
DROP INDEX IF EXISTS uniq_active_subscription;
DROP INDEX IF EXISTS idx_transactions_subscription_created;



