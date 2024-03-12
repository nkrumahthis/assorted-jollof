<?php

namespace App\Http\Controllers;

use App\Models\Payment;

class PaymentController extends Controller
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

        return Payment::create([

        ]);
    }

    //
    public function index()
    {
        return Payment::all();
    }

    public function show($id)
    {
        return Payment::find($id);
    }
}
