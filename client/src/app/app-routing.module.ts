import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { ShowComponent } from './components/show/show.component';
import { LoginComponent } from './components/login/login.component';
import { AuthGuard } from './auth/auth.guard';

const routes: Routes = [
  { path: 'login', component: LoginComponent },
  {
    path: 'show',
    component: ShowComponent,
    canActivate: [AuthGuard]
  },
  {
    path: '',
    redirectTo: '/login',
    pathMatch: 'full'
  },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
