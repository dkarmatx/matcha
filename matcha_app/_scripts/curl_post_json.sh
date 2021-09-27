curl    -X POST \
        -H "Content-type: application/json" \
        --data "@$2" \
        --silent \
        "http://localhost:8080$1" |

json_pp