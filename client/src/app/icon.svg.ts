import { MatIconRegistry } from "@angular/material";
import { DomSanitizer } from "@angular/platform-browser";
//这个文件用来定义访问图标用的url
export const loadSvgsources = (ir: MatIconRegistry, ds: DomSanitizer) => {
    ir.addSvgIcon('arrow_down', ds.bypassSecurityTrustResourceUrl(`assets/icon/arrow_down.svg`));
    ir.addSvgIcon('arrow_right', ds.bypassSecurityTrustResourceUrl(`assets/icon/arrow_right.svg`));
    ir.addSvgIcon('log', ds.bypassSecurityTrustResourceUrl(`assets/icon/log.svg`));
}