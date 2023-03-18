import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

// Add these imports at the top of your time.service.ts file
import { interval } from 'rxjs';
import { switchMap } from 'rxjs/operators';

@Injectable({
  providedIn: 'root',
})
export class TimeService {
  private readonly API_URL = 'http://localhost:8080/time';

  constructor(private http: HttpClient) { }

  // Modify the getCurrentTime() method in your TimeService
  getCurrentTimeAutoRefresh(intervalMs: number): Observable<string> {
    return interval(intervalMs).pipe(
      switchMap(() => this.http.get(this.API_URL, { responseType: 'text' }))
    );
  }

}
