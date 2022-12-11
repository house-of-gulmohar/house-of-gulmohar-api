import { Global, Module } from '@nestjs/common';
import { Pool } from 'pg';
import { CONNECTION_POOL } from 'src/utils/constants';
import {
  ConfigurableDatabaseModule,
  DATABASE_OPTIONS,
} from './database.module-definition';
import { DatabaseService } from './database.service';
import { DatabaseOptions } from './interfaces/databaseOptions.interface';

// Initializing connection pool and providing as a object
const DatabaseProvider = {
  provide: CONNECTION_POOL,
  inject: [DATABASE_OPTIONS],
  useFactory: (databaseOptions: DatabaseOptions) => {
    return new Pool({
      host: databaseOptions.host,
      port: databaseOptions.port,
      user: databaseOptions.user,
      password: databaseOptions.password,
      database: databaseOptions.database,
      connectionTimeoutMillis: 5000,
    });
  },
};

@Global()
@Module({
  providers: [DatabaseProvider, DatabaseService],
  exports: [DatabaseService],
})
export class DatabaseModule extends ConfigurableDatabaseModule {}
