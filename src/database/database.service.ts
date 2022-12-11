import { Inject, Injectable } from '@nestjs/common';
import { Pool } from 'pg';
import { CONNECTION_POOL } from 'src/utils/constants';

@Injectable()
export class DatabaseService {
  constructor(@Inject(CONNECTION_POOL) private pool: Pool) {
    this.pool
      .connect()
      .then(() => {
        console.log('connection success');
      })
      .catch((err) => {
        console.log(err);
        console.log('connection failed');
      });
  }
}
