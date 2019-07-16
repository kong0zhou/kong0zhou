<?php 
//设置最大执行时间
set_time_limit(0);
function getHtml($url){
  // 1. 初始化
   $ch = curl_init();
   // 2. 设置选项，包括URL

   //设置需要获取的url地址
   curl_setopt($ch,CURLOPT_URL,$url);
   //设置是否以字符串形式返回
   curl_setopt($ch,CURLOPT_RETURNTRANSFER,1);
   //
   curl_setopt($ch,CURLOPT_HEADER,0);
   // 3. 执行并获取HTML文档内容
   $output = curl_exec($ch);
   if($output === FALSE ){
    $output = '';
   }
   // 4. 释放curl句柄
   curl_close($ch);
   return $output;
}
function getPageData($url){
  // 获取整个网页内容
  $html = getHtml($url);
  $coding = mb_detect_encoding($html, array("ASCII","GB2312","GBK","UTF-8"));  
  if ($coding != "UTF-8" || !mb_check_encoding($html, "UTF-8"))  $html = mb_convert_encoding($html, 'utf-8', 'GBK,UTF-8,ASCII');
  $pattern ='/https?:\/\/[^,"\'<\n]*/';
  preg_match_all($pattern,$html,$result_array);
  return $result_array;
}

$arr = getPageData("https://tieba.baidu.com/");
foreach($arr[0] as $i=>$value){
    echo $i;
    echo "    ";
    echo $value;
    echo "\n";
}

?>
