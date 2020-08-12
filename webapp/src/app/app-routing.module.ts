import {NgModule} from '@angular/core';
import {Routes, RouterModule} from '@angular/router';
import {CarComponent} from './components/car/car.component';
import {KontaktsComponent} from './components/kontakts/kontakts.component';
import {CardComponent} from './components/card/card.component';

const routes: Routes = [
    {path: '', component: CarComponent},
    {path: 'about', component: KontaktsComponent},
    {path: 'card', component: CardComponent},
];

@NgModule({
    imports: [RouterModule.forRoot(routes)],
    exports: [RouterModule]
})
export class AppRoutingModule {
}
