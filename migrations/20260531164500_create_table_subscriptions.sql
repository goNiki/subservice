-- +goose Up
CREATE TABLE subscriptions (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    service_name VARCHAR(255) NOT NULL,
    status VARCHAR(20) NOT NULL,
    current_period_start DATE NOT NULL,
    current_period_end DATE NOT NULL,
    created_at TIMESTAMPTZ NOT NULL default NOW(),
    updated_at TIMESTAMPTZ NOT NULL default NOW(), 
    

    CHECK (
        status IN (
            'active',
            'expired',
            'canceled',
            'paused'
        )
    ) 
);



-- +goose Down
DROP TABLE IF EXISTS subscriptions;
