import React from "react";
import {MdDelete} from "react-icons/md";

const MovieTable = ({movies, updateFunction}) => {
  const deleteMovie = async id => {
    const deleteMovie = await fetch(`http://localhost:8080/movies/${id}`, {
      method: "DELETE",
      headers: {
        "Content-Type": "application/json"
      }
    });
    const data = await deleteMovie.json();
    console.log(`Movie with Id:${id} deleted !`);
    console.log(data);
    updateFunction();
  };

  return (
    <div className='bg-green-100 w-1/2'>
      {movies.length !== 0 && (
        <table>
          <thead>
            <tr>
              <th>ID</th>
              <th>Title</th>
              <th>ISBN</th>
              <th>Director</th>
            </tr>
          </thead>
          <tbody>
            {movies.length !== 0 &&
              movies.map(item => (
                <tr key={item.id}>
                  <td>{item.id}</td>
                  <td>{item.title}</td>
                  <td>{item.isbn}</td>
                  <td>
                    {item.director &&
                      `${item.director.firstname} ${item.director.lastName}`}
                  </td>
                  <td>
                    <MdDelete
                      fontSize={28}
                      className='cursor-pointer'
                      onClick={e => {
                        deleteMovie(item.id);
                      }}
                    />
                  </td>
                </tr>
              ))}
          </tbody>
        </table>
      )}
    </div>
  );
};

export default MovieTable;
