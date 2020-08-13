import {Component, OnInit} from '@angular/core';
import {Hero} from '../heroes/hero';
import {HeroService} from '../../services/hero.service';

@Component({
    selector: 'app-hero-dashboard',
    templateUrl: './hero-dashboard.component.html',
    styleUrls: ['./hero-dashboard.component.css']
})
export class HeroDashboardComponent implements OnInit {
    heroes: Hero[] = [];

    constructor(private heroService: HeroService) {
    }

    ngOnInit(): void {
        this.getHeroes();
    }

    private getHeroes(): void {
        this.heroService.getHeroes().subscribe(heroes => this.heroes = heroes.slice(1, 5));
    }
}
