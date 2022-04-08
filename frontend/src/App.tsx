import { VFC } from 'react';
import { BrowserRouter, Link, Route, Routes } from 'react-router-dom';

import { Admin } from './components/Admin';
import { Categories } from './components/Categories';
import { CategoryPage } from './components/CategoryPage';
import { Home } from './components/Home';
import { OneMovie } from './components/OneMovie';
import { Movies } from './components/Movies';

export const App: VFC = () => {
  return (
    <BrowserRouter>
      <div className="container">
        <div className="row">
          <h1 className="mt-3">Go Watch a Movie!</h1>
          <hr className="mb-3" />
        </div>

        <div className="row">
          <div className="col-md-2">
            <nav>
              <ul className="list-group">
                <li className="list-group-item">
                  <Link to="/">Home</Link>
                </li>
                <li className="list-group-item">
                  <Link to="/movies">Movies</Link>
                </li>
                <li className="list-group-item">
                  <Link to="/by-category">Categories</Link>
                </li>
                <li className="list-group-item">
                  <Link to="/admin">Manage Catalogue</Link>
                </li>
              </ul>
            </nav>
          </div>
          <div className="col-md-10">
            <Routes>
              <Route path="/movies/:id" element={<OneMovie />} />
              <Route path="/movies" element={<Movies />} />
              <Route path="/by-category" element={<CategoryPage />} />
              <Route path="/by-category/drama" element={<Categories title={`Drama`} />} />
              <Route path="/by-category/comedy" element={<Categories title={`Comedy`} />} />
              <Route path="/admin" element={<Admin />} />
              <Route path="/" element={<Home />} />
            </Routes>
          </div>
        </div>
      </div>
    </BrowserRouter>
  );
};
