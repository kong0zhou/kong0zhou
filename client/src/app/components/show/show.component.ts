import { Component, OnInit, ElementRef, ViewChild } from '@angular/core';
import { NestedTreeControl } from '@angular/cdk/tree';
import { MainService, FileNode } from "../../services/main.service"
import { MatTreeNestedDataSource } from '@angular/material/tree';
import { ReplyProto, ReqProto } from "../../msg-proto";
import { BehaviorSubject } from 'rxjs';
import { ToastrService } from 'ngx-toastr';

import { LogHighLightPipe } from '../../pipe/log-high-light.pipe'

// import { default as AnsiUp } from 'ansi_up'

@Component({
  selector: 'app-show',
  templateUrl: './show.component.html',
  styleUrls: ['./show.component.css']
})
export class ShowComponent implements OnInit {
  @ViewChild('cutOff', { static: true }) cutOff: ElementRef;
  @ViewChild('main', { static: true }) mainDiv: ElementRef;
  @ViewChild('right', { static: true }) right: ElementRef;

  constructor(
    public service: MainService,
    private toastr: ToastrService,
  ) { }


  // 当ctrlB==2快捷键成立
  ctrlB: number = 0;

  // >>>>>>>>>>>>>>  动态css  >>>>>>>>>>>>>>>>>
  // 左边div的宽度，单位是%
  leftWidth: number = 10

  leftWidthStyleMethod() {
    let leftWidthStyle = {
      'width': this.leftWidth + '%'
    }
    return leftWidthStyle
  }

  cutOffLeftMethod() {
    let cutOffLeft = {
      'left': this.leftWidth + '%'
    }
    return cutOffLeft
  }

  rightWidthStyleMethod() {
    let rightWidth = {
      'width': (100 - this.leftWidth) + '%'
    }
    return rightWidth
  }
  //<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<

  // >>>>>>>>>>>> 文件树 >>>>>>>>>>>
  nestedTreeControl: NestedTreeControl<FileNode>;
  nestedDataSource: MatTreeNestedDataSource<FileNode>;
  hasNestedChild = (_: number, nodeData: FileNode) => !nodeData.isFile;
  getChildren = (node: FileNode) => node.children;

  expand(node) {
    this.nestedTreeControl.expand(node)
    // console.log('click')
  }

  ngOnInit() {
    this.nestedTreeControl = new NestedTreeControl<FileNode>(this.getChildren);
    this.nestedDataSource = new MatTreeNestedDataSource();

    //<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<

    // >>>>>>>> 捕捉ctrl+B的事件  >>>>>>>>>>>>>
    document.onkeydown = (event) => {
      if (event.keyCode == 17) {
        this.ctrlB++;
        // console.log(this.ctrlB)
      }
      else if (event.keyCode == 66 && this.ctrlB == 1) {
        this.ctrlB++;
        // console.log('在这里执行ctrl+B的函数')
        this.lefthiddenChange()
        this.ctrlB--
        // console.log(this.ctrlB)
      }
      else {
        this.ctrlB = 0
        // console.log(this.ctrlB)
      }
    }
    document.onkeyup = (event) => {
      if (event.keyCode == 17) {
        this.ctrlB = 0;
        // console.log(this.ctrlB)
      }
    }
    // <<<<<<<<<<<<<<<<<<<<<<<<<<

    // >>>>>>>> 捕捉鼠标拖动cutOff事件，并实现侧边栏的放大缩小  >>>>>>>>>>>
    let leftDiv = <HTMLDivElement>(this.cutOff.nativeElement)
    let mainD = <HTMLDivElement>(this.mainDiv.nativeElement)
    leftDiv.onmousedown = (e) => {
      document.onmousemove = (ei) => {
        // console.log(ei.screenX)
        let templeftWidth = (ei.clientX / mainD.clientWidth) * 100
        // 左边div设置个范围 8% ～ 60%
        if (templeftWidth > 8 && templeftWidth < 60) {
          this.leftWidth = templeftWidth
        }
      }
    }
    document.onmouseup = (e) => {
      // console.log('stop')
      document.onmousemove = null
    }
    // <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
    // >>>>>>>>>>>>>>> 从后端获取所有的文件名和路径 >>>>>>>>>>>>>>>>

    this.service.getAllFile().subscribe(
      (data) => {
        if (data.status == 0) {
          // this.allfiles = data.data
          // console.log(JSON.stringify(this.service.listToTree(data.data)))
          this.nestedDataSource.data = this.service.listToTree(data.data)
        } else {
          console.error(data.msg)
          this.toastr.error(data.msg, "错误提示")
        }
      },
      (err) => {
        console.error(err)
      }
    )

    // <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
  }

  //>>>>>>>>>>>隐藏左边div事件>>>>>>>>>>>>>>>>>>>>>>>
  isLeftHidden: boolean = false
  // 暂时存储左边div隐藏前的div的宽度
  tempLeftWidth: number
  lefthiddenChange() {
    this.isLeftHidden = !this.isLeftHidden
    if (this.isLeftHidden) {
      this.tempLeftWidth = this.leftWidth
      this.leftWidth = 0
    } else {
      this.leftWidth = this.tempLeftWidth
    }
  }

  nowFile: string = ''
  selectFile(file: string) {
    this.nowFile = file
  }
  // <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<

  // >>>>>>>>>>>>>>获得log文件的内容>>>>>>>>>>>>>>>>
  nowFilePath: string = ''
  nowFileName: string = '';

  logText: string = '';
  logHTML: string = '';

  canScroll: boolean = false;

  logTextPipe: LogHighLightPipe = new LogHighLightPipe;
  clickFile(node: FileNode) {
    if (node.filePath == '' || typeof node.filePath == 'undefined' || node.filePath == null) {
      console.error('node.filePath is null or undefined')
      return
    }
    // console.log(node.filePath)
    // 检查是否已经选中了这个node
    if (node.filePath == this.nowFilePath) {
      return
    } else {
      this.nowFilePath = node.filePath
    }

    this.nowFileName = node.filename;
    if (this.service.source != null) {
      this.service.sseClose()
      this.logText = ''
      this.logHTML = ''
    }

    let right = <HTMLDivElement>(this.right.nativeElement)
    this.service.getFileText(node.filePath).subscribe(
      (data) => {
        // right.scrollTop = right.scrollHeight
        if (right.scrollTop + right.clientHeight <= right.scrollHeight + 2 && right.scrollTop + right.clientHeight >= right.scrollHeight - 2) {
          this.canScroll = true;
          // console.log('在底部')
        } else {
          this.canScroll = false;
        }
        let reply = JSON.parse(data)
        // console.log(reply.data)
        this.logText = this.logText + reply.data
        this.logHTML = this.logHTML + this.logTextPipe.transform(reply.data) + '<br>';
        right.innerHTML = this.logHTML;
        // right div滚动条自动滚到底部
        if (this.canScroll) {
          right.scrollTop = right.scrollHeight
        }
      },
      (error) => {
        console.error(error)
        this.service.sseClose()
        this.logText = ''
        this.logHTML = ''
      }
    )
  }
  // <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
}
