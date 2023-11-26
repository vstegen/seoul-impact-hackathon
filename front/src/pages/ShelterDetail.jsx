import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { PaperClipIcon } from "@heroicons/react/20/solid";
import axios from "axios";
import IconImg from "../assets/images/chlidren_2.png";
import Map from "../assets/images/maps.png";

export default function ShelterDetail() {
  const [shelter, setShelter] = useState({});
  const { id } = useParams();
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await axios.get(
          `http://localhost:8000/shelters/${id}`
        );
        console.log("response-" + JSON.stringify(response));
        setShelter(response.data);
        setLoading(false);
      } catch (error) {
        console.log(error);
      }
    };
    fetchData();
  }, [id]);

  return (
    <div className="mx-auto max-w-7xl px-6 pb-32 pt-36 sm:pt-30 lg:px-8 lg:pt-32">
      <div className="overflow-hidden bg-white shadow sm:rounded-lg">
        <div className="px-4 py-6 sm:px-6">
          <p class="text-2xl ...">{shelter.name}</p>
          <br />
          {shelter.currentStatus == "Open" ? (
            <span class="bg-green-100 text-red-800 text-xs font-medium me-2 px-2.5 py-0.5 rounded-full dark:bg-green-900 dark:text-green-300">
              {shelter.currentStatus}
            </span>
          ) : (
            <span class="bg-red-100 text-red-800 text-xs font-medium me-2 px-2.5 py-0.5 rounded-full dark:bg-red-900 dark:text-red-300">
              {shelter.currentStatus}
            </span>
          )}
          <br />
          <h5>
            {shelter.rating}
            <br />
            {shelter.capacity} / {shelter.maxCapacity} spots filled
            <br />
            <br />
          </h5>
          <p class="text-sm ...">{shelter.facilities}</p>
          <br />
          <p className="mt-1 max-w-2xl text-sm leading-6 text-gray-500">
            {shelter.announcement}
          </p>
          <img className="w-30 bg-gray-300" src={IconImg} alt="" />
        </div>
        <div className="border-t border-gray-100">
          <dl className="divide-y divide-gray-100">
            <div className="px-4 py-6 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
              <dt className="text-sm font-medium text-gray-900">Address</dt>
              <dd className="mt-1 text-sm leading-6 text-gray-700 sm:col-span-2 sm:mt-0">
                {shelter.address}
              </dd>
            </div>
            <div className="px-4 py-6 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
              <dt className="text-sm font-medium text-gray-900">OpeningTime</dt>
              <dd className="mt-1 text-sm leading-6 text-gray-700 sm:col-span-2 sm:mt-0">
                {shelter.openingTime}
              </dd>
            </div>
            <div className="px-4 py-6 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
              <dt className="text-sm font-medium text-gray-900">Contact</dt>
              <dd className="mt-1 text-sm leading-6 text-gray-700 sm:col-span-2 sm:mt-0">
                {shelter.contact}
              </dd>
            </div>
            <div className="px-4 py-6 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
              <dt className="text-sm font-medium text-gray-900">Website</dt>
              <dd className="mt-1 text-sm leading-6 text-gray-700 sm:col-span-2 sm:mt-0">
                {shelter.website}
              </dd>
            </div>
            <div className="px-4 py-6 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
              <dt className="text-sm font-medium text-gray-900">Desc</dt>
              <dd className="mt-1 text-sm leading-6 text-gray-700 sm:col-span-2 sm:mt-0">
                {!shelter.desc ? "-" : shelter.desc}
              </dd>
            </div>
            <div className="px-4 py-6 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
              <dt className="text-sm font-medium text-gray-900">
                Requirements
              </dt>
              <dd className="mt-1 text-sm leading-6 text-gray-700 sm:col-span-2 sm:mt-0">
                {!shelter.requirements ? "-" : shelter.requirements}
              </dd>
            </div>
            <div className="px-4 py-6 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6">
              <dt className="text-sm font-medium leading-6 text-gray-900">
                Map
              </dt>
              <dd className="mt-2 text-sm text-gray-900 sm:col-span-2 sm:mt-0">
                <img src={Map} />
              </dd>
            </div>
          </dl>
        </div>
      </div>
    </div>
  );
}
