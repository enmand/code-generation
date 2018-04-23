<?php
$console = require('console');

// Do some setup for a Node environment
if(!function_exists('printf')) {
    function printf($msg) {
        $console->log($msg);
    }
}

class phpConsole {
    function log($msg) {
        printf("%s\n", $msg);
    }
}

if(!isset($console) || $console === 1) {
    $console = new phpConsole();
}

function consolelog($msg) {
    $console->log(msg);
}


// Code fun begins here
$ar = array('this', 'is', 'an', 'array', 'of', 'things', 1, true);
foreach($ar as $elem) {
    $console->log(gettype($elem));
}

class Something {
    private $x = "hello";
    public $y = "world";

    function __constructor($x = null, $y = null) {
        if(!is_null($x)) {
            $this->x = $x;
        }

        if(!is_null($y)) {
            $this->y = $y;
        }
    }

    function getX() {
        return $this->x;
    }
}

$s = new Something(1, 2);
$console->log($s->getX() . " " . $s->y);
