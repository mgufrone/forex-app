
-- +migrate Up
-- +migrate StatementBegin
CREATE TABLE rates (
    id varchar(255) not null primary key,
    base varchar(3) not null,
    symbol varchar(3) not null,
    source varchar(20) not null,
    source_type varchar(20) not null,
    sell decimal(10, 2) not null,
    buy decimal(10, 2) not null,
    updated_at datetime not null
);
CREATE INDEX idx_rate_base on rate (base);
CREATE INDEX idx_rate_symbol on rate (symbol);
CREATE INDEX idx_rate_source on rate (`source`);
CREATE INDEX idx_rate_source_type on rate (source_type);
CREATE INDEX idx_rate_updated_at on rate (updated_at desc);
-- +migrate StatementEnd


-- +migrate Down
drop table rate;
