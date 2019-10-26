import { Injectable } from '@angular/core';
// import { Observable } from 'rxjs';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { ReplyProto, ReqProto } from "../msg-proto";
import { Observable, of } from 'rxjs';
import { environment } from '../../environments/environment'

@Injectable({
  providedIn: 'root'
})
export class MainService {

  api = environment.api

  constructor(
    public http: HttpClient,
  ) { }

  getAllFile() {
    return this.http.get<ReplyProto>(this.api + '/apiAllFile')
  }

  source: EventSource
  isReallyClosed: boolean = false
  getFileText(filePath: string) {
    // source.readyState
    //  0 — connecting
    //  1 — open
    //  2 — closed

    if (!EventSource) {
      console.error("EventSource unsupported");
      return;
    }
    if (filePath == null || filePath == '' || typeof filePath == 'undefined') {
      console.error('filePath is null or empty')
      return
    }
    let req: ReqProto = {
      data: filePath
    }
    this.source = new EventSource(this.api + '/apiShow?q=' + JSON.stringify(req))
    //每隔30s执行一次
    // setInterval(() => {
    //   if (this.source.readyState !== 1 && !this.isReallyClosed) {
    //     console.log("进行重连")
    //     this.source.close()

    //     this.source = new EventSource(this.api + '/apiShow?q=' + JSON.stringify(req))
    //   }
    // }, 30 * 1000)
    return new Observable<string>(
      observer => {
        this.source.onopen = () => {
          console.log('sse通道已建立…')
        }
        this.source.onmessage = (event) => {
          console.log('onmessage')
          observer.next(event.data);
        }
        this.source.onerror = (event) => {
          // if (this.source.readyState !== 1){
          //   return
          // }
          console.error(event)
          observer.error(event);
        };
        this.source.addEventListener("isconnected", (event) => {
          console.log("确定sse建立")
        })
        // this.source.addEventListener("nop",(event)=>{
        //   console.log("接受心跳包")
        // })
      }
    )
  }



  sseClose() {
    this.source.close();
    this.isReallyClosed = true;
  }

  // 这个函数的每个文件路径最终都应指向一个文件
  listToTree(arr: string[]) {
    if (arr == null || arr.length == 0) {
      console.error("arr is null or empty")
      return
    }
    let ret: FileNode[] = [];
    for (let i = 0; i < arr.length; ++i) {
      let path = arr[i].split("/");
      let _ret = ret;
      for (let j = 0; j < path.length; ++j) {
        let filename = path[j];
        let obj: FileNode = null;
        for (let k = 0; k < _ret.length; ++k) {
          let _obj = _ret[k];
          if (_obj.filename === filename) {
            obj = _obj;
            break;
          }
        }
        if (!obj) {
          obj = new FileNode;
          obj.filename = filename;
          if (j != path.length - 1) {
            obj.children = []
            obj.isFile = false
          } else {
            obj.isFile = true
            obj.children = null
          }
          _ret.push(obj);
        }
        if (obj.children) _ret = obj.children;
        else obj.filePath = arr[i]
      }
    }
    return ret;
  }

  login(uid: string, password: string) {
    let req: ReqProto = {
      data: {
        uid: uid,
        password: password,
      }
    }
    return this.http.put<ReplyProto>(this.api + '/apiLogin', req)
  }

  checkUser() {
    return this.http.get<ReplyProto>(this.api + '/apiCheckUser')
  }
}

export class FileNode {
  children: FileNode[];
  filename: string;
  isFile: boolean;
  // 如果是文件，会存储它的文件路径
  filePath: string
}
