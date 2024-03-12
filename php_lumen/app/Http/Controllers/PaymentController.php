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
        $validated = $request->validate([
            "amount" => 'required|float',
            'customer_id' => 'required|integer',
            'order_id' => 'required|integer'
        ]);

        return Payment::create([
            'amount' => $request->amount,
            'customer_id' => $request->customer_id,
            'order_id' => $request->order_id
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
