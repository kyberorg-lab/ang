import {Injectable} from '@angular/core';
import {Hero} from '../components/heroes/hero';
import {Observable, of} from 'rxjs';
import {MessageService} from './message.service';
import {HttpClient, HttpResponse} from '@angular/common/http';

@Injectable({
    providedIn: 'root'
})
export class HeroService {
    private heroesUrl = 'http://localhost:8080/api/heroes';
    private heroes: any = [];
    private hero: Hero;

    constructor(private messageService: MessageService, private httpClient: HttpClient) {
    }

    getAllHeroes(): Observable<HttpResponse<Hero[]>> {
        return this.httpClient.get<Hero[]>(this.heroesUrl, {observe: 'response'});
    }

    getHeroes(): Observable<Hero[]> {
        this.messageService.add('HeroService: fetched heroes');
        this.getAllHeroes().subscribe(resp => {
            for (const data of resp.body) {
                this.heroes.push(data);
            }
        });

        return this.heroes;
    }

    getHero(id: number): Observable<Hero> {
        this.messageService.add(`HeroService: fetched hero id=${id}`);

        this.httpClient.get<Hero>(this.heroesUrl + '/' + id, {observe: 'response'})
            .subscribe(resp => this.hero = resp.body);

        return of(this.hero);
    }
}
