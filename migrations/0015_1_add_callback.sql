-- +migrate Up
alter table application
	add column callback text;

-- +migrate Down
alter table application
	drop column callback;
