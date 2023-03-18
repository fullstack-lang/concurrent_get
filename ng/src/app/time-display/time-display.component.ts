import { Component, OnInit, OnDestroy } from '@angular/core';
import { TimeService } from '../time.service';
import { Subscription } from 'rxjs';

@Component({
  selector: 'app-time-display',
  templateUrl: './time-display.component.html',
  styleUrls: ['./time-display.component.css']
})
export class TimeDisplayComponent implements OnInit, OnDestroy {
  currentTime: string = ""
  private timeSubscription: Subscription = new Subscription

  constructor(private timeService: TimeService) { }

  ngOnInit(): void {
    this.startAutoRefresh(500); // Refresh every 500 ms (half second)
  }

  ngOnDestroy(): void {
    this.stopAutoRefresh();
  }

  startAutoRefresh(intervalMs: number): void {
    this.timeSubscription = this.timeService
      .getCurrentTimeAutoRefresh(intervalMs)
      .subscribe((time: string) => {
        this.currentTime = time;
      });
  }

  stopAutoRefresh(): void {
    if (this.timeSubscription) {
      this.timeSubscription.unsubscribe();
    }
  }
}
