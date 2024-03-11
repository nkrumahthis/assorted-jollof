<?php

/** @var \Laravel\Lumen\Routing\Router $router */

/*
|--------------------------------------------------------------------------
| Application Routes
|--------------------------------------------------------------------------
|
| Here is where you can register all of the routes for an application.
| It is a breeze. Simply tell Lumen the URIs it should respond to
| and give it the Closure to call when that URI is requested.
|
*/

$router->get('/', function () use ($router) {
    return $router->app->version();
});

$router->group(["prefix" => "/api/v1"], function () use ($router) {
    $router->get("users", "UserController@index");
    $router->get("users/{id}", "UserController@show");

    $router->post("orders", "OrderController@create");
    $router->get("orders", "OrderController@index");
    $router->get("orders/{id}", "OrderController@show");

    $router->post("customers", "CustomerController@create");
    $router->get("customers", "CustomerController@index");
    $router->get("customers/{id}", "CustomerController@show");

    $router->post("payments", "PaymentController@create");
    $router->get("payments", "PaymentController@index");
    $router->get("payments/{id}", "PaymentController@show");

});
