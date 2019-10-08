import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
  name: 'logHighLight'
})
export class LogHighLightPipe implements PipeTransform {

  transform(value: any, ...args: any[]): any {
    return null;
  }

}
