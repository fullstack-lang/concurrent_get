// Modify the imports at the top of your app.component.ts file
import { Component, OnInit, OnDestroy } from '@angular/core';
import { TimeService } from './time.service';
import { Subscription } from 'rxjs';

// Modify the AppComponent class
@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  instances = Array(200).fill(0);
}
