import {Response} from '../../core/entities/response';

export class PrimeRepository {
    async findLargestPrimeUnder(n: number): Promise<Response> {
        const serviceUrl = `${process.env.REACT_APP_PRIME_SERVICE_URL}/${process.env.REACT_APP_PRIME_SERVICE_ENDPOINT}/${n}`;
        
        const res = await fetch(serviceUrl);
        const jsonData = await res.json();
        return new Response(jsonData.errorCode, jsonData.message, jsonData.value);
    }
}