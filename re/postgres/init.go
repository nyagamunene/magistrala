// Copyright (c) Abstract Machines
// SPDX-License-Identifier: Apache-2.0

package postgres

import (
	_ "github.com/jackc/pgx/v5/stdlib" // required for SQL access
	migrate "github.com/rubenv/sql-migrate"
)

func Migration() *migrate.MemoryMigrationSource {
	return &migrate.MemoryMigrationSource{
		Migrations: []*migrate.Migration{
			{
				Id: "rules_01",
				// VARCHAR(36) for colums with IDs as UUIDS have a maximum of 36 characters
				// STATUS 0 to imply enabled and 1 to imply disabled
				Up: []string{
					`CREATE TABLE IF NOT EXISTS rules (
						id                VARCHAR(36) PRIMARY KEY,
						name              VARCHAR(1024),
						domain_id         VARCHAR(36) NOT NULL,
						metadata          JSONB,
						created_by        VARCHAR(254),
						created_at        TIMESTAMP,
						updated_at        TIMESTAMP,
						updated_by        VARCHAR(254),
						input_channel     VARCHAR(36),
						input_topic       TEXT,
						output_channel    VARCHAR(36),
						output_topic      TEXT,
						status            SMALLINT NOT NULL DEFAULT 0 CHECK (status >= 0),
						logic_type        SMALLINT NOT NULL DEFAULT 0 CHECK (status >= 0),
						logic_value       BYTEA,
						time              TIMESTAMP,
						recurring         SMALLINT,
						recurring_period  SMALLINT,
						start_datetime    TIMESTAMP
					)`,
				},
				Down: []string{
					`DROP TABLE IF EXISTS rules`,
				},
			},
			{
				Id: "rules_02",
				Up: []string{
					`CREATE TABLE IF NOT EXISTS report_config (
						id          	 	VARCHAR(36) PRIMARY KEY,
						name				VARCHAR(1024),
						domain_id         	VARCHAR(36) NOT NULL,
						"limit"				BIGINT CHECK ("limit" >= 0),
						channel_ids 	  	TEXT[],
						client_ids 		  	TEXT[],
						metrics				TEXT[],
						"to" 				TEXT[],
						"from" 				TEXT,
						subject 			TEXT,
						status				SMALLINT NOT NULL DEFAULT 0 CHECK (status >= 0),
						created_at			TIMESTAMP,
						created_by			VARCHAR(254),
						updated_at			TIMESTAMP,
						updated_by			VARCHAR(254),
						time              	TIMESTAMP,
						recurring         	SMALLINT,
						recurring_period  	SMALLINT,
						start_datetime    	TIMESTAMP,
						config			  	JSONB
					);`,
				},
				Down: []string{
					`DROP TABLE IF EXISTS report_config;`,
				},
			},
		},
	}
}
