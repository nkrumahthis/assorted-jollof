<?php

namespace App\Http\Controllers;

use App\Models\Order;

class OrderController extends Controller
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
            'packs' => 'required|integer',
            'customerId' => 'required|integer',
            'location' => 'required|string',
            'status' => 'required|string'
        ]);

        return Order::create([
            'packs' => $request->packs,
            'customerId' => $request->customerId,
            'location' => $request->location,
            'status' => $request->status
        ]);
    }

    //
    public function index()
    {
        return Order::all();
    }

    public function show($id)
    {
        return Order::find($id);
    }
}
