#!/usr/bin/env php
<?php

class Translate
{
    private $_httpContext = null;

    function __construct()
    {
        $url = parse_url($_ENV['http_proxy']);
        $options = array(
            'http' => array(
                'proxy' => "tcp://{$url['host']}:{$url['port']}",
                'request_fulluri' => true)
            );

        if (is_null($url['host'])) {
            unset($options['http']);
        }

        $this->_httpContext = stream_context_create($options);
    }

    function google($args, $to = 'ja')
    {
        $baseURI = 'http://ajax.googleapis.com/ajax/services/language';

        $src = urlencode(implode($args, ' '));
        $detect = json_decode($this->file_get_contents("$baseURI/detect?v=1.0&q={$src}"));
        $from = $detect->responseData->language;
        $to = ($to == $from) ? 'en' : $to;

        $langPair = "langpair={$from}%7C{$to}";
        $result = json_decode($this->file_get_contents("$baseURI/translate?v=1.0&q={$src}&{$langPair}"));

        return $result->responseData->translatedText;
    }

    function file_get_contents($uri)
    {
        return file_get_contents($uri, false, $this->_httpContext);
    }
}

$cmd = array_shift($argv);
$translate = new Translate();
echo $translate->google($argv);
echo "\n";