import React, { useEffect, useState } from "react";
import axios from "axios";
import IconImg from "../assets/images/heart_2.png";
import { Link } from "react-router-dom";

export default function Restaurant() {
  const [restaurants, setRestaurants] = useState([]);
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await axios.get("http://localhost:8000/food");
        // console.log("response-" + JSON.stringify(response));
        setRestaurants(response.data);
        setLoading(false);
      } catch (error) {
        console.log(error);
      }
    };
    fetchData();
  }, []);

  return (
    <>
      <div className="overflow-hidden">
        <div className="mx-auto max-w-7xl px-6 pb-32 pt-36 sm:pt-30 lg:px-8 lg:pt-32">
          <ul role="list" className="divide-y divide-gray-100">
            {restaurants?.map((restaurant, idx) => (
              <Link to={`/restaurant/${restaurant.id}`} key={idx}>
                <li className="flex justify-between gap-x-6 py-5">
                  <div className="flex min-w-0 gap-x-4">
                    <img
                      className="h-12 w-12 flex-none rounded-full bg-gray-50"
                      src={IconImg}
                      alt=""
                    />
                    <div className="min-w-0 flex-auto">
                      <p className="text-sm font-semibold leading-6 text-gray-900">
                        {restaurant.name}
                      </p>
                      <p className="mt-1 truncate text-xs leading-5 text-gray-500">
                        {restaurant.contact}
                      </p>
                      <p className="mt-1 truncate text-xs leading-5 text-gray-500">
                        {restaurant.website}
                      </p>
                    </div>
                  </div>
                  <div className="hidden shrink-0 sm:flex sm:flex-col sm:items-end">
                    <p className="text-sm leading-6 text-gray-900">
                      {restaurant.foodTypes}
                    </p>
                    {restaurant.currentStatus == "Open" ? (
                      <div className="mt-1 flex items-center gap-x-1.5">
                        <div className="flex-none rounded-full bg-emerald-500/20 p-1">
                          <div className="h-1.5 w-1.5 rounded-full bg-emerald-500" />
                        </div>
                        <p className="text-xs leading-5 text-gray-500">
                          {restaurant.currentStatus}
                        </p>
                      </div>
                    ) : (
                      <div className="mt-1 flex items-center gap-x-1.5">
                        <div className="flex-none rounded-full bg-red-500/20 p-1">
                          <div className="h-1.5 w-1.5 rounded-full bg-red-500" />
                        </div>
                        <p className="text-xs leading-5 text-gray-500">
                          {restaurant.currentStatus}
                        </p>
                      </div>
                    )}
                  </div>
                </li>
              </Link>
            ))}
          </ul>
        </div>
      </div>
    </>
  );
}
