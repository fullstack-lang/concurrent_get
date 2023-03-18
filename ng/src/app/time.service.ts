import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

// Add these imports at the top of your time.service.ts file
import { interval } from 'rxjs';
import { switchMap } from 'rxjs/operators';
import { map } from 'rxjs/operators';


export interface TimeResponse {
  UnixMilli: number;
}

@Injectable({
  providedIn: 'root',
})
export class TimeService {
  private readonly API_URL = 'http://localhost:8080/time';

  constructor(private http: HttpClient) { }

  getCurrentTimeAutoRefresh(intervalMs: number): Observable<number> {
    return interval(intervalMs).pipe(
      switchMap(() => this.http.get<TimeResponse>(this.API_URL)),
      map((response) => response.UnixMilli)
    );
  }

}
