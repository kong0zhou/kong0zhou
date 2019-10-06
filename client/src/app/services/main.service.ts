import { Injectable } from '@angular/core';
// import { Observable } from 'rxjs';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { ReplyProto, ReqProto } from "../msg-proto";
import { Observable, of } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class MainService {


  constructor(
    public http: HttpClient,
  ) { }

  getAllFile() {
    return this.http.get<ReplyProto>('/log-api/allFile')
  }
  listToTree(arr: string[]) {
    let ret:FileNode[]= [];
    for (let i = 0; i < arr.length; ++i) {
      let path = arr[i].split("/");
      let _ret = ret;
      for (let j = 0; j < path.length; ++j) {
        let filename = path[j];
        let obj:FileNode = null;
        for (let k = 0; k < _ret.length; ++k) {
          let _obj = _ret[k];
          if (_obj.filename === filename) {
            obj = _obj;
            break;
          }
        }
        if (!obj) {
          obj=new FileNode;
          obj.filename = filename;
          if (filename.indexOf(".") < 0) {
            obj.children = []
            obj.isFile=false
          }else {
            obj.isFile=true
            obj.children=null
          }
          _ret.push(obj);
        }
        if (obj.children) _ret = obj.children;
      }
    }
    return ret;
  }
}

export class FileNode {
  children: FileNode[];
  filename: string;
  isFile: boolean;
}
