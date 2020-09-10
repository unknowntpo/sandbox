<?php
$mystring = 'absolute string';
$findme = 'ab';
$pos = strpos($mystring, $findme);

if ($pos === FALSE)
  echo "The string '$findme' does not found\n";
else
  echo "The string '$findme' is found at position '$pos'\n";

var_dump($pos);
?>
