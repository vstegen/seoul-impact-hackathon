import React from "react";
import { Routes, Route } from "react-router-dom";

import Home from "./pages/Home";
import About from "./pages/About";
import Contact from "./pages/Contact";
import Restaurant from "./pages/Restaurant";
import Shelter from "./pages/Shelter";
import Support from "./pages/Support";
import ShelterDetail from "./pages/ShelterDetail";
import RestaurantDetail from "./pages/RestaurantDetail";

const Router = () => {
  return (
    <Routes>
      <Route path="/" element={<Home />} />
      <Route path="/about" element={<About />} />
      <Route path="/contact" element={<Contact />} />
      <Route path="/restaurants" element={<Restaurant />} />
      <Route path="/restaurant/:id" element={<RestaurantDetail />} />
      <Route path="/shelters" element={<Shelter />} />
      <Route path="/shelter/:id" element={<ShelterDetail />} />
      <Route path="/support" element={<Support />} />
    </Routes>
  );
};

export default Router;
