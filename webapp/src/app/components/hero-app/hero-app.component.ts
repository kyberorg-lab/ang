import {Component, OnInit} from '@angular/core';

@Component({
    selector: 'app-hero-app',
    templateUrl: './hero-app.component.html',
    styleUrls: ['./hero-app.component.css']
})
export class HeroAppComponent implements OnInit {
    title = 'Tour of Heroes';

    constructor() {
    }

    ngOnInit(): void {
    }

}
