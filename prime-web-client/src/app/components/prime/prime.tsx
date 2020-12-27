import {Fragment, useState} from 'react';
import {PrimeUsecase} from '../../../core/usecases/prime.usecase';
import './prime.css';

function Prime() {
    const [n, setNumber] = useState(0);
    const [prime, setPrime] = useState(0);

    const handleNumberChanged = (e: React.FormEvent<HTMLInputElement>) => {
        const input = e.currentTarget.value;
        const inputNumber = parseInt(input);

        setNumber(inputNumber);
    }

    const calculatePrimeNUmber = async () => {
        const primeService = new PrimeUsecase();
        const response = await primeService.findLargestPrimeUnder(n);

        if (response.errorCode === 0) {
            setPrime(response.value);
        }
    }

    return (
        <Fragment>
            <div className="prime-text-container-shape prime-text-container-presentation">
                <input type="text" value = {n === 0 || Number.isNaN(n) ? '' : n} onChange={handleNumberChanged} placeholder="Enter number" className="prime-text-shape prime-text-presentation"/>
            </div>
            <br/>
            <div>
                <span>{prime === 0 || Number.isNaN(prime) ? '----' : prime} </span>
            </div>
            <br/>
            <button onClick={calculatePrimeNUmber} className="prime-button-shape prime-button-presentation">
                Calculate
            </button>
        </Fragment>
    );
  }
  
  export default Prime;