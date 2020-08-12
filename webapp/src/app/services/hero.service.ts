import {Injectable} from '@angular/core';
import {Hero} from '../components/heroes/hero';
import {HEROES} from '../components/heroes/mock-heroes';

@Injectable({
    providedIn: 'root'
})
export class HeroService {

    constructor() {
    }

    getHeroes(): Hero[] {
        return HEROES;
    }
}
