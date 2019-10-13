import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ErrorStateMatcher } from '@angular/material/core';
import { MainService, FileNode } from "../../services/main.service"
import { FormControl, FormGroupDirective, NgForm, Validators } from '@angular/forms';
import { ToastrService } from 'ngx-toastr';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  uid: string = '';
  password: string;

  constructor(
    public service:MainService,
    private router: Router,
    private toastr: ToastrService,
  ) { }
  // >>>>>>>> 验证器 >>>>>>>>>>
  uidFormControl = new FormControl('', [
    Validators.required,
    
  ])

  passwordFormControl = new FormControl('', [
    Validators.required,
  ])
  // <<<<<<<<<<<<<<<<<<<<<

  // >>>>>>>>>动态css>>>>>>>>>>
  mtop: number;
  formDivStyle() {
    let form = {
      'margin-top': this.mtop + 'px'
    }
    return form;
  }
  // <<<<<<<<<<<<<<<<<<

  ngOnInit() {
    this.mtop = (document.documentElement.clientHeight / 2) - 120 - 50
  }

  onKeyup(value, type) {
    if (typeof value == 'undefined' || value == null || value == '') {
      console.error("error:value is null")
    }
    if (typeof type == 'undefined' || type == null || value == '') {
      console.error("error:type is null")
    }
    switch (type) {
      case 'uid':
        this.uid = value
        break;
      case 'password':
        this.password = value;
        break;
    }
  }

  login(){
    this.service.login(this.uid, this.password).subscribe(res =>{
      // console.log(res)
      if (res.status==0){
        this.router.navigate(["/show"])
      }else{
          console.error(res.msg)
          this.toastr.error(res.msg,"错误提示")
      }
    },
    (error)=>{
      console.error(error)
    })
  }
}
