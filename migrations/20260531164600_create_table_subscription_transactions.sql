-- +goose Up
CREATE TABLE subscription_transactions (
    id UUID PRIMARY KEY,
    subscription_id UUID NOT NULL REFERENCES subscriptions(id) ON DELETE CASCADE,
    transaction_type VARCHAR(20) NOT NULL,
    price INTEGER NOT NULL CHECK(price > 0),
    period_start DATE NOT NULL,
    period_end DATE NOT NULL,
    created_at TIMESTAMPTZ NOT NULL default NOW(),

    

    CHECK (
        transaction_type IN (
            'created',
            'renewed',
            'canceled'
        )
    ) 
);



-- +goose Down
DROP TABLE IF EXISTS subscription_transactions;
