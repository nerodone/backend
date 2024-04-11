#!/bin/fish


source ./.env.fish

goose -dir ./sql/schema/ postgres $XATA_PG down-to 0

goose -dir ./sql/schema/ postgres $XATA_PG up
