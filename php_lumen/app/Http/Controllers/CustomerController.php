<?php

namespace App\Http\Controllers;

use App\Models\Customer;

class CustomerController extends Controller
{
    /**
     * Create a new controller instance.
     *
     * @return void
     */
    public function __construct()
    {
        //
    }

    public function create($request)
    {
        $this->validate($request, [

        ]);

        return Customer::create([

        ]);
    }

    //
    public function index()
    {
        return Customer::all();
    }

    public function show($id)
    {
        return Customer::find($id);
    }
}
