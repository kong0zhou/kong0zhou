import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
  name: 'logHighLight'
})
export class LogHighLightPipe implements PipeTransform {

  transform(value: string, ...args: any[]): any {
    let reg:RegExp=new RegExp("\\[E\\]",'ig');
    value=value.replace(reg,'<span style="color: red">[E]</span>')
    return value
  }

}
