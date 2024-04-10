#!/usr/bin/bash

echo >migration.sql

find ./sql/schema/ -name "*.sql" -print0 | sort -z | xargs -0 -I {} sh -c 'less "{}" | sed -n "/-- +goose Up/,/-- +goose Down/p" | sed "s/-- +goose.*//" >>migration.sql'
