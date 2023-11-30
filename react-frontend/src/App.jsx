import React, {useEffect, useState} from "react";
import "./App.css";
import MovieById from "./components/MovieById";
import AddMovie from "./components/AddMovie";

function App() {
  const [movies, setMovies] = useState([]);
  const [resData, setResData] = useState();

  const getMovies = async () => {
    const res = await fetch("http://localhost:8080/");
    const data = await res?.json();
    console.log(data);
    setMovies(data);
  };

  useEffect(() => {
    getMovies();
  }, []);

  return (
    <div className='App'>
      <h2>CRUD API GO</h2>
      <div className='flex flex-col justify-center gap-y-10'>
        <div className='flex justify-between p-4 mx-3 mt-5 border border-black gap-x-3'>
          <div className='bg-purple-100 w-1/2'>
            {" "}
            <MovieById setData={setResData} />
            <AddMovie />
          </div>

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
                      </tr>
                    ))}
                </tbody>
              </table>
            )}
          </div>
        </div>
        <div className='mx-auto'>{resData && JSON.stringify(resData)}</div>
      </div>
    </div>
  );
}

export default App;
