import {Injectable} from '@angular/core';
import {Hero} from '../components/heroes/hero';
import {Observable, of} from 'rxjs';
import {MessageService} from './message.service';
import {HttpClient, HttpHeaders, HttpParams} from '@angular/common/http';
import {catchError, tap} from 'rxjs/operators';

@Injectable({
    providedIn: 'root'
})
export class HeroService {
    private heroesUrl = 'http://localhost:8080/api/heroes';

    private headers = new HttpHeaders({'Content-Type': 'application/json'});
    private httpOptions = {headers: this.headers};

    constructor(private messageService: MessageService, private httpClient: HttpClient) {
    }

    getHeroes(): Observable<Hero[]> {
        return this.httpClient
            .get<Hero[]>(this.heroesUrl, this.httpOptions)
            .pipe(
                tap(_ => this.log('fetched heroes')),
                catchError(this.handleError<Hero[]>('getHeroes', []))
            );
    }

    getHero(id: number): Observable<Hero> {
        const httpParams = new HttpParams().set('id', id + '');

        return this.httpClient.get<Hero>(this.heroesUrl,
            {headers: this.headers, params: httpParams})
            .pipe(
                tap(_ => this.log(`fetched hero id=${id}`)),
                catchError(this.handleError<Hero>(`getHero id=${id}`))
            );

    }

    updateHero(hero: Hero): Observable<any> {
        return this.httpClient.put(this.heroesUrl, hero, this.httpOptions)
            .pipe(
                tap(_ => this.log(`updated hero id=${hero.id}`)),
                catchError(this.handleError<any>('updateHero'))
            );
    }

    addHero(hero: Hero): Observable<Hero> {
        return this.httpClient.post<Hero>(this.heroesUrl, hero, this.httpOptions)
            .pipe(
                tap((newHero: Hero) => this.log(`added hero w/ id=${newHero.id}`)),
                catchError(this.handleError<Hero>('addHero'))
            );
    }

    deleteHero(hero: Hero): Observable<Hero> {
        const id = typeof hero === 'number' ? hero : hero.id;
        const httpParams = new HttpParams().set('id', id + '');

        return this.httpClient.delete<Hero>(this.heroesUrl, {headers: this.headers, params: httpParams})
            .pipe(
                tap(_ => this.log(`deleted hero id=${id}`)),
                catchError(this.handleError<Hero>('deleteHero'))
            );
    }

    private log(message: string): void {
        this.messageService.add(`HeroService: ${message}`);
    }

    private handleError<T>(operation = 'operation', result?: T): any {
        return (error: any): Observable<T> => {
            console.error(error);

            this.log(`${operation} failed: ${error.message}`);

            return of(result as T);
        };
    }

}
