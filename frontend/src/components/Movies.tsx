import { memo, useEffect, useState, VFC } from 'react';
import { Link } from 'react-router-dom';

import { Movie } from '../types/movie';

export const Movies: VFC = memo(() => {
  const [movies, setMovies] = useState<Movie[]>([]);

  useEffect(() => {
    setMovies([
      { id: 1, title: 'The Shawshank Redemption', runtime: 142 },
      { id: 2, title: 'The Godfather', runtime: 175 },
      { id: 3, title: 'The Dark Knight', runtime: 153 },
    ]);
  }, []);

  return (
    <>
      <h2>Choose a movie</h2>
      <ul>
        {movies.map((movie) => (
          <li key={movie.id}>
            <Link to={`/movies/${movie.id}`}>{movie.title}</Link>
          </li>
        ))}
      </ul>
    </>
  );
});
