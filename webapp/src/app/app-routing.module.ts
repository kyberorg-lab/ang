import {NgModule} from '@angular/core';
import {Routes, RouterModule} from '@angular/router';
import {CarComponent} from './components/car/car.component';
import {IndexComponent} from './components/index/index.component';
import {HeroesComponent} from './components/heroes/heroes.component';

const routes: Routes = [
    {path: '', component: IndexComponent},
    {path: 'cars', component: CarComponent},
    {path: 'heroes', component: HeroesComponent},
];

@NgModule({
    imports: [RouterModule.forRoot(routes)],
    exports: [RouterModule]
})
export class AppRoutingModule {
}
