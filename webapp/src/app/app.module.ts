import {BrowserModule} from '@angular/platform-browser';
import {NgModule} from '@angular/core';

import {AppRoutingModule} from './app-routing.module';
import {AppComponent} from './app.component';
import {CarComponent} from './components/car/car.component';
import {FormsModule} from '@angular/forms';
import {KontaktsComponent} from './components/kontakts/kontakts.component';
import {CardComponent} from './components/card/card.component';
import {FormComponent} from './components/form/form.component';

@NgModule({
    declarations: [
        AppComponent,
        CarComponent,
        KontaktsComponent,
        CardComponent,
        FormComponent
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
