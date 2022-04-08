import { memo, useEffect, useState, VFC } from 'react';
import { useParams } from 'react-router-dom';

import { Movie } from '../types/movie';

export const OneMovie: VFC = memo(() => {
  const { id } = useParams();

  const [movie, setMovie] = useState<Movie>();

  useEffect(() => {
    setMovie({
      id: Number(id),
      title: 'Some movie',
      runtime: 150,
    });
  }, []);

  return (
    <>
      <h2>
        Movie: {movie?.title} {movie?.id}
      </h2>
      <table className="table table-compact table-striped">
        <thead></thead>
        <tbody>
          <tr>
            <td>
              <strong>Title:</strong>
            </td>
            <td>{movie?.title}</td>
          </tr>
          <tr>
            <td>
              <strong>Run time:</strong>
            </td>
            <td>{movie?.runtime} minutes</td>
          </tr>
        </tbody>
      </table>
    </>
  );
});
