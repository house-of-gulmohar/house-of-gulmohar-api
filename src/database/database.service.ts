import { Inject, Injectable, Logger, Type } from '@nestjs/common';
import { Pool } from 'pg';
import { from, Observable } from 'rxjs';
import { CONNECTION_POOL } from 'src/utils/constants';

@Injectable()
export class DatabaseService {
  logger = new Logger(DatabaseService.name);
  constructor(@Inject(CONNECTION_POOL) private pool: Pool) {
    this.pool
      .connect()
      .then(() => {
        this.logger.log('connected to postgres');
      })
      .catch((err) => {
        this.logger.error('db connection failed');
        this.logger.error(err);
      });
  }

  private runQuery(query: string, params: any[]): Observable<any> {
    const now = Date.now();
    return from(this.pool.query(query, params)).pipe();
  }
}
