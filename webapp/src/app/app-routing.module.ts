import {NgModule} from '@angular/core';
import {RouterModule, Routes} from '@angular/router';
import {CarComponent} from './components/car/car.component';
import {IndexComponent} from './components/index/index.component';
import {HeroDashboardComponent} from './components/hero-dashboard/hero-dashboard.component';
import {HeroesComponent} from './components/heroes/heroes.component';
import {HeroDetailComponent} from './components/hero-detail/hero-detail.component';
import {HeroAppComponent} from './components/hero-app/hero-app.component';
import {NotesComponent} from './components/notes/notes.component';

const routes: Routes = [
    {path: '', component: IndexComponent},
    {path: 'cars', component: CarComponent},
    {
        path: 'heroes', component: HeroAppComponent, children: [
            {path: 'list', component: HeroesComponent},
            {path: 'dashboard', component: HeroDashboardComponent},
            {path: 'details/:id', component: HeroDetailComponent},
            {path: '', redirectTo: 'dashboard', pathMatch: 'full'},
        ]
    },
    {path: 'notes', component: NotesComponent},
];

@NgModule({
    imports: [RouterModule.forRoot(routes)],
    exports: [RouterModule]
})
export class AppRoutingModule {
}
