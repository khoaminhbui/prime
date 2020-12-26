import {Response} from '../../core/entities/response';

export class PrimeRepository {
    async findLargestPrimeUnder(n: number): Promise<Response> {
        const res = await fetch(`http://localhost:8000/prime/${n}`);
        const jsonData = await res.json();
        return new Response(jsonData.errorCode, jsonData.message, jsonData.value);
    }
}