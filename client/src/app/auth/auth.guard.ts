import { Injectable } from '@angular/core';
import { CanActivate, ActivatedRouteSnapshot, RouterStateSnapshot, UrlTree } from '@angular/router';
import { Observable, of } from 'rxjs';
import { MainService, FileNode } from "../services/main.service"
import { ToastrService } from 'ngx-toastr';
import { Router } from '@angular/router';
import { map, concatMap, catchError } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class AuthGuard implements CanActivate {
  constructor(
    private service: MainService,
    private router: Router,
    private toastr: ToastrService,
  ) { }

  canActivate(
    next: ActivatedRouteSnapshot,
    state: RouterStateSnapshot): Observable<boolean> {
    return this.checkUser();
  }
  checkUser(): Observable<boolean> {
    return this.service.checkUser().pipe(map(res => {
      if (res.status == 0) {
        return true
      }else{
        this.router.navigate(['../login']);
        return false
      }
    },error=>{
      console.error(error)
      this.router.navigate(['../login']);
      return false
    }
    ))
  }
}

