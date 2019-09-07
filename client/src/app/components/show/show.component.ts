import { Component, OnInit, ElementRef } from '@angular/core';
import { MainService } from "../../services/main.service"
import { ReplyProto, ReqProto } from "../../msg-proto";

@Component({
  selector: 'app-show',
  templateUrl: './show.component.html',
  styleUrls: ['./show.component.css']
})
export class ShowComponent implements OnInit {

  fileData: ReplyProto[]
  logText: string = ''

  constructor(
    public main: MainService,
    public el: ElementRef
  ) { }

  ngOnInit() {
    console.log(document.documentElement.clientHeight)
    this.el.nativeElement.querySelector('.main').style.height = (document.documentElement.clientHeight-20) +'px';
    this.main.createObservableSocket("ws://localhost:8083/show").subscribe(
      data => {
        this.logText = this.logText + '<br>' + <string>data
        // this.el.nativeElement.querySelector('.right').style.height = document.body.clientHeight;
        this.el.nativeElement.querySelector('.right').scrollTop = this.el.nativeElement.querySelector('.right').scrollHeight
      },
      err => console.error(err),
      () => console.log("流已经结束")
    )
    this.main.getAllFile().subscribe(
      data => {
        this.fileData = data.data
      },
      err => console.error(err)
    )
  }
  // send(){
  //   this.main.sendMessage("sdfgt")
  // }

  nowPath: string
  getLogFile(p: string) {
    if (this.nowPath == p) {
      return
    }
    this.nowPath = p
    let req: ReqProto = {
      data: p,
    }
    this.logText = ""
    this.main.sendMessage(req)
  }
}
