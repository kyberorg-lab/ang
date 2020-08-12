import {Component, OnInit} from '@angular/core';

@Component({
    selector: 'app-car',
    templateUrl: './car.component.html',
    styleUrls: ['./car.component.css']
})
export class CarComponent implements OnInit {

    name: string;
    speed: number;
    model: string;
    colors: Colors;
    engine: {
        volume: number,
        fuel: string,
    };
    options: string[];

    constructor() {
    }

    ngOnInit(): void {
        this.name = 'UAZ';
        this.speed = 90;
        this.model = 'Patriot';
        this.colors = {
            car: 'Orange',
            salon: 'Grey',
            wheels: 'Black'
        };
        this.engine = {
            volume: 2.7,
            fuel: 'bensin'
        };
        this.options = ['ABS', '4WD'];
    }
}

interface Colors {
    car: string;
    salon: string;
    wheels: string;
}
