module Main exposing (..)


main =
    Platform.worker
        { init = init
        , update = update
        , subscriptions = subscriptions
        }


type alias Flags =
    { initial : Int }


type alias Model =
    Int


type Msg
    = Greeting Int


init : Flags -> ( Model, Cmd Msg )
init flags =
    ( flags.initial, Cmd.none )


update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
    case msg of
        Greeting int ->
            ( model + int, Cmd.none )


subscriptions : Model -> Sub Msg
subscriptions model =
    Sub.none
