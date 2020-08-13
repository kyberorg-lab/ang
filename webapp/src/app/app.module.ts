import {BrowserModule} from '@angular/platform-browser';
import {NgModule} from '@angular/core';

import {AppRoutingModule} from './app-routing.module';
import {AppComponent} from './app.component';
import {CarComponent} from './components/car/car.component';
import {FormsModule} from '@angular/forms';
import {IndexComponent} from './components/index/index.component';
import {HeroesComponent} from './components/heroes/heroes.component';
import {HeroDetailComponent} from './components/hero-detail/hero-detail.component';
import {MessagesComponent} from './components/messages/messages.component';
import {HeroDashboardComponent} from './components/hero-dashboard/hero-dashboard.component';
import {HeroAppComponent} from './components/hero-app/hero-app.component';

@NgModule({
    declarations: [
        AppComponent,
        CarComponent,
        IndexComponent,
        HeroesComponent,
        HeroDetailComponent,
        MessagesComponent,
        HeroDashboardComponent,
        HeroAppComponent,
    ],
    imports: [
        BrowserModule,
        AppRoutingModule,
        FormsModule
    ],
    providers: [],
    bootstrap: [AppComponent]
})
export class AppModule {
}
